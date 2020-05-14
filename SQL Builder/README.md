# BIM SQL Builder

Comprehensive guide also located in the PCH Wiki.

BIM SQL builder scripts are used to convert information stored in an Excel style format (.xlsx or similar) to an SQL INSERT script style. There are three types.

## SQL_Builder.py

Information required to be updated in the database must be copied into the Template.xlsx file from the Source Points List workbook. Python script that accepts two arguments:

Source File path (must be /Template.xlsx format):

```python
#Loads excel file - must be in the correct format
bimData = load_workbook('.\\SQL Builder\\Source Excel Docs\\Assorted.xlsx')
```

Output File:

```python
#Opens blank sql file
with open('.\\SQL Builder\\Output SQL\\DCE UPS 2019.sql','w') as f:
```

## SQLBuilderStandalone.py

Information required to be updated in the database must be copied into the Template.xlsx file from the Source Points List workbook. Python script does not have any argument. It will convert any files within the same directory as the script. Can convert multiple Template files.


## SQL_Room_Builder.py

This script looks directly at the ELV-Combined 2018.xlsx document and accepts the path as an argument:

```python
roomData = load_workbook('E:\\Perth Childrens Hospital\\BIM\\Room Information\\ELV-RM-Combined 2018.xlsx')
```

It creates an INSERT script into a file as an argument:
```python
f = open('E:\\Code Projects\\SQL Builder\\Output SQL\\dbo_room.sql','w')
```