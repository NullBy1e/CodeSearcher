import re

file_csv = open("formatted_results.csv","w")
file_csv.write("File_Location, Line\n")

with open("results.txt","rt") as file:
    results = file.read()
    lines = results.split("\n")
    for line in lines:
        formatted_line = re.sub('[\t,\n,]','',line)
        try:
            values = formatted_line.split("/;#;/")
            file_csv.write(values[1]+","+values[0]+"\n")
        except IndexError:
            print("Invalid line: ",formatted_line)

file_csv.close()