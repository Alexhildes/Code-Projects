#Author Alex Hildebrand
#May 2020
#The height of the Covid 19 Pandemic....

import time
import sys
from pyModbusTCP.client import ModbusClient
from pyModbusTCP import utils

# Create the InfluxDB client object
client = InfluxDBClient(host='localhost', port=8086, database='cnReach')

#config
host_ip = "192.168.15.191"
port_client = 502

#cnReach Details
clients =   {
            cnReach1 : {
                address : '192.168.15.190',
                location : 'Head Office'
                radio : 'Access Point'
            cnReach2 : {
                address : '192.168.15.191',
                location : 'Workshop'
                radio : 'End Point'
            cnReach3 : {
                address : '192.168.15.192',
                location : 'Pumping Station 1'
                radio : 'Repeater'
            cnReach4 : {
                address : '192.168.15.193'
                location : 'Pumping Station 2'
                radio : 'End Point'
            }

#Registers
input_registers =   {
                    uptime : 2
                    voltage : 4
                    temp : 6
                    tamper : 
                    }

discrete_registers =    {
                        tamper
                        }

#Host connect
c = ModbusClient(host_ip, port=port_client, auto_open=True, debug=False)

#2 x 16 Bit Register to Float Conversion. Big Endian must be FALSE
def read_float(address, number=1):
    reg_l = c.read_input_registers(address, number * 2)
    if reg_l:
        return [utils.decode_ieee(f) for f in utils.word_list_to_long(reg_l, False)]
    else:
        return None

def write_data(measurement, location, temperature, voltage, uptime)
     # Create the JSON data structure
        data = [{
            "measurement": measurement,
                "tags": {
                    "location": location,
                    "radio" : radio,
                },
                "fields": {
                    "temperature" : temperature,
                    "voltage" : voltage,
                    "uptime" : uptime
                }
            }]
        # Send the JSON data to InfluxDB
        client.write_points(data)
        
        print(iso + ' - Data Sent to InfluxDB')
#Loop
while True:
    try:                
    
        #Looking at Register 10008 which is PLC Address 10008, Protocol Address 7
        regs = c.read_discrete_inputs(7)        #Digital Input 8 Modbus Register 7
        regs.append(read_float(4))
        if regs:
            if regs[0] == True:
                print('Light is on!')
            else:
                print('Light is off :(')
            print(regs[1])

        else:
            print("Error reading things!")
        
        time.sleep(5)
    except KeyboardInterrupt:
        print("Exiting")
        sys.exit(0)

    except Exception as e:
        print("Received exception!")
        print(e)
        time.sleep(10)