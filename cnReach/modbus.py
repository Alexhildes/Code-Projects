import time
from pyModbusTCP.client import ModbusClient
from pyModbusTCP import utils

#Sample code from cnReach
#DR7=fnDRGet(7);
#PRINT DR7;
#
#IF DR7=1 THEN
#    fnHRSet(272, 20000);
#    PRINT "Door Open";
#ELSE
#    fnHRSet(272,0)
#    PRINT "Door Closed"
#ENDIF

#config
host_ip = "192.168.15.191"
port_client = 502

#Host connect
c = ModbusClient(host_ip, port=port_client, auto_open=True, debug=False)


def read_float(self, address, number=1):
    reg_l = self.read_holding_registers(address, number * 2)
    if reg_l:
        return [utils.decode_ieee(f) for f in utils.word_list_to_long(reg_l)]
    else:
        return None

while True:
    try:                
        

        #Looking at Register 10008 which is PLC Address 10008, Protocol Address 7
        regs = c.read_discrete_inputs(7)
        regs.append(c.read_input_registers(7,2))
        regs.append(c.read_input_registers(8))
        regs.append(read_float(7))
        if regs:
            if regs[0] == True:
                #print(regs[0])
                print('Light is on!')
            else:

                print('Light is off :(')
            print(regs[1])
            print(regs[2])

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