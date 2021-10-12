import re

print("Starting")
file_csv = open("formatted_results.csv","w")
file_csv.write("File_Location, Line\n")
with open("results.txt","rt") as file:
    results = file.read()
    lines = results.split("\n")
    for line in lines:
        # * Remove tabs and new lines
        for char in '\t':
            formatted_line = line.replace(char,'')
        for char in '\n':
            formatted_line = line.replace(char,'')
        formatted_line = re.sub('[\t,\n,]','',line)
        # * Write to formatted_results
        try:
            values = formatted_line.split("/;#;/")
            file_csv.write(values[1]+","+values[0]+"\n")
        except IndexError:
            print("Invalid line: ",formatted_line)

file_csv.close()
print("Finished")