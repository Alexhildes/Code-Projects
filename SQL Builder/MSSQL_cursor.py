from os import getenv
import pymssql
with pymssql.connect(server='localhost',database='bim') as conn:
    cursor = conn.cursor()
    cursor.execute('select * from device where id between 10000 and 20000')
    #conn.commit()
    results = cursor.fetchone()
    print(results)

