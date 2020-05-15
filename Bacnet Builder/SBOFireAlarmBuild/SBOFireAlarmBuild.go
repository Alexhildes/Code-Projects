// bacnetbuilder
package main

import (
    "text/template"
    "flag"
    "fmt"
    "os"
    "encoding/csv"
    "io"
    "bytes"
    "io/ioutil"
)

type AlarmEntry struct {
    Name    string
	Description  string
	AlarmMessage     string
	ResetMessage     string
	Notification string
	RefObject    string 
	Category     string
	AlarmValue   string
	EventEnable string
}

const sboExportTmpl = `
	<OI NAME="{{.Name}}" DESCR="{{.Description}}" TYPE="bacnet.EventEnrollment">
      <PI Name="AlarmMessage" Value="{{.AlarmMessage}}" />
      <PI Name="BACnetName" Value="{{.Name}}" />
      <PI Name="Category">
        <Reference DeltaFilter="0" Locked="1" Object="~/System/Alarm Control Panel/Alarm Handling/Categories/{{.Category}}" Retransmit="0" TransferRate="10" />
      </PI>
      <PI Name="EventEnable" Value="{{EventEnable}}" />
      <PI Name="EventType" Value="1" />
      <PI Name="NotificationClass">
        <Reference DeltaFilter="0" Object="{{.Notification}}" Retransmit="0" TransferRate="10" />
      </PI>
      <PI Name="ObjectPropertyReference">
        <Reference DeltaFilter="0" Object="../{{.RefObject}}" Property="Value" Retransmit="0" TransferRate="10" />
      </PI>
      <PI Name="ResetMessage" Value="{{.ResetMessage}}" />
      <OI NAME="Change of state event parameters" TYPE="bacnet.pt.eventparameters.COS" hidden="1">
        <OI NAME="bacnet.pt.listofstates.ListOfUnsigned32" TYPE="bacnet.pt.listofstates.ListOfUnsigned32" hidden="1">
          <OI NAME="bacnet.pt.listofstates.ObjectOfUnsignedValue" TYPE="bacnet.pt.listofstates.ObjectOfUnsignedValue" hidden="1">
            <PI Name="Value" Value="{{.AlarmValue}}" />
          </OI>
        </OI>
      </OI>
    </OI>	
	
`

func main() {
    var filename string
    var outfilename string
    var skipHeader bool

    flag.StringVar(&filename, "f", "points.csv", "specify the points list csv file.  Defaults to points.csv")
    flag.StringVar(&outfilename, "o", "points.xml", "specify the output export filename.  Defaults to points.xml")
    flag.BoolVar(&skipHeader, "h", true, "specify if the points csv file has a header.  Defaults to true")
    flag.Parse()

    alarms, err := parsePoints(filename, skipHeader)

    if err != nil {
        panic(err)
    }

    buf := &bytes.Buffer{}
    t := template.Must(template.New(sboExportTmpl).Parse(sboExportTmpl))

    buf.WriteString(`<?xml version="1.0" encoding="utf-8"?>
<ObjectSet ExportMode="Standard" Version="1.6.1.5000" Note="TypesFirst">
  <ExportedObjects>`)


    for i:=0; i < len(alarms); i++ {
        alarm := alarms[i]
        if err := t.Execute(buf, alarm); err != nil {
            panic(err)
        }
    }
    buf.WriteString(`
  </ExportedObjects>
</ObjectSet>
`)


    err = ioutil.WriteFile(outfilename, buf.Bytes(), 0644)
    if err != nil {
        panic(err)
    }

    fmt.Println("Done.")
}

func parsePoints(filename string, skipHeader bool) ([]AlarmEntry, error) {
    f, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    csvr := csv.NewReader(f)

    points := []AlarmEntry{}

    if skipHeader {
        _,_ = csvr.Read()
    }

    for {
        row, err := csvr.Read()
        if err != nil {
            if err==io.EOF {
                err = nil
            }
            return points, err
        }

        ae := AlarmEntry{}

        name := row[0]
		desc := row[4]
		alarmMsg := row[5]
		rstMsg := row[6]
		category := row[7]
		priority := row[8]
		eventType := row[9]
		refObj := row[11]
		almVal := row[12]
		eventenable := row[13]
		
		ae.Name = name
		ae.Description = desc
		ae.AlarmMessage = alarmMsg
		ae.ResetMessage = rstMsg
		ae.Notification = "../../_Notification/" + eventType + "_" + priority
		ae.RefObject = refObj		
		ae.Category = category
		ae.AlarmValue = almVal
		ae.EventEnable = eventenable

        points = append(points, ae)
    }
    return points, nil
}