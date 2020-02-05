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
    Name         string
    Description  string
    BacNetName   string    
	BoundObject  string
	BoundProperty    string
}

const sboExportTmpl = `
    <OI DESCR="{{.Description}}" NAME="{{.Name}}" TYPE="bacnet.point.multistate.Value">
      <PI Name="BACnetName" Value="{{.BacNetName}}" />
      <PI Name="EventEnable" Value="0" />
	  <PI Name="Value" Type="public.aaauqfpf4dvhg6cletmtybsxhj_aaauqfpf4dvhg6cletmqavicm6">
        <Reference DeltaFilter="0" Object="{{.BoundObject}}" Property="{{.BoundProperty}}" Retransmit="0" TransferRate="10" />
      </PI>	  
      <PI Name="NumberOfStates" Value="5" />
      <PI Name="StateText" Value="Normal;Isolate;Fault;Pre-Alarm;Alarm;" />
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
		loop := row[10]
		

        ae.Name = name
		ae.Description = desc
        ae.BacNetName = name
  		ae.BoundObject = "../../_Programs/ImportLoop" + loop
		ae.BoundProperty = "Out" + name		


        points = append(points, ae)
    }
    return points, nil
}