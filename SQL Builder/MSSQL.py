## From SQL to DataFrame Pandas
 import pandas as pd
 import pyodbc
 sql_conn = pyodbc.connect("DRIVER={ODBC Driver 13 for SQL Server};"
                             "SERVER=ip;"
                             "DATABASE=dbname;"
                             #"Trusted_Connection=yes"
                             "uid=sa;pwd=password") 
 query="physical sql query"
 df=pd.read_sql(query,sql_conn)
 print(df)'
