#Author Alex Hildebrand
#May 2020
#The height of the Covid 19 Pandemic....

import time
import sys
from pyModbusTCP.client import ModbusClient
from pyModbusTCP import utils
from influxdb import InfluxDBClient

# Create the InfluxDB client object
client = InfluxDBClient(host='localhost', port=8086, database='cnReach')

#Modbus
port_client = 502

#cnReach Details
clients =   {
            "cnReach1" : {
                "ipaddress" : "192.168.15.190",
                "location" : "Head Office",
                "radio" : "End Point",
            },
            "cnReach2" : {
                "ipaddress" : "192.168.15.191",
                "location" : "Workshop",
                "radio" : "Access Point",
            },
            "cnReach3" : {
                "ipaddress" : "192.168.15.192",
                "location" : "Pumping Station 1",
                "radio" : "Repeater",
            },
            "cnReach4" : {
                "ipaddress" : "192.168.15.193",
                "location" : "Pumping Station 2",
                "radio" : "End Point",
                    }
            }       

#Registers
input_registers =   {
                    'uptime' : 2,
                    'voltage' : 4,
                    'temp' : 6
                    }

discrete_registers =    {
                        'tamper' : 7
                        }



#2 x 16 Bit Register to Float Conversion. Big Endian must be FALSE
def read_float(client, address, number=1):

    #Connnect to each Client
    c = ModbusClient(client, port=port_client, auto_open=True, debug=False)

    reg_l = c.read_input_registers(address, number * 2)
    if reg_l:
        return [utils.decode_ieee(f) for f in utils.word_list_to_long(reg_l, False)]
    else:
        return None

def read_state(client, address):

    #Connnect to each Client
    c = ModbusClient(client, port=port_client, auto_open=True, debug=False)

    reg = c.read_discrete_inputs(address)
    
    if reg:
        return None
    else:
        return None
        
def write_data(measurement, location, device, radio, temperature, voltage, uptime):
     # Create the JSON data structure
        iso = time.ctime()
        data = [{
            "measurement": measurement,
                "tags": {
                    "device" : device,
                    "location": location,
                    "radio" : radio,
                },
                "fields": {
                    "temperature" : temperature,
                    "voltage" : voltage,
                    "uptime" : uptime,
                }
            }]
        # Send the JSON data to InfluxDB
        client.write_points(data)
        print(iso + ' - Data Sent to InfluxDB')

def write_states(measurement, location, device, radio, message):
     # Create the JSON data structure
    iso = time.ctime()

    #This will be a mini alarm handler
    data = [{
            "measurement": measurement,
                "tags": {
                    "device" : device,
                    "location": location,
                    "radio" : radio,
                },
                "fields": {
                    "message" : message,
                }
            }]
#Loop
while True:
    try:
        #Reading floats from the Registers and sending them to the Database
        #k is keys, v is value                
        for k, v in clients.items():
            
            try:
                write_data("tblRadioMetrics",
                            v['location'],
                            k,
                            v['radio'],
                            read_float(v['ipaddress'],input_registers['temp'])[0],
                            read_float(v['ipaddress'],input_registers['voltage'])[0],
                            read_float(v['ipaddress'],input_registers['uptime'])[0],
                            )
            except Exception:
                print('Exception with ' + k)
                pass

        #Reading discretes from the Registers and sending them to the database
        #try:
            

            #write_states("tblRadioAlarms",
            #            clients["CnReach4"]['location'],
            #            "cnReach4",
            #            clients["CnReach4"]['radio'],
            #            alarm_message)

            #)

        time.sleep(30)
    except KeyboardInterrupt:
        print("Exiting")
        sys.exit(0)

    except Exception as e:
        print("Received exception!")
        print(e)
        time.sleep(10)