import sys,os, time,hashlib
from hashlib import md5

initalListOfFiles = []
filesToWatch = dict()
userRequestedFile = ""
targetFileExtension = ".sav"
secondsBetweenRefreshes = 5
archiveFolderName = "backupPilots"

#simplify file name retrieval
def files(path):
    for file in os.listdir(path):
        if os.path.isfile(os.path.join(path, file)):
            yield file

#get the initial list of files
for file in files("."):
    initalListOfFiles.append(file)

# All files, or just a specific file?
if(len(userRequestedFile)<=0):
    print("No specific file to watch was requested.")
    print("Watching all *"+targetFileExtension+" files in this folder.")
    for file in initalListOfFiles:
        if (targetFileExtension in file):
            filesToWatch[file] = ''
else:
    for file in initalListOfFiles:
        if(file == userRequestedFile):
            filesToWatch.append(file)

print(" ")

#bailout if nothing to watch
if(len(filesToWatch) <= 0):
    print("Found no files to watch.")
    sys.exit(0)


#seed proper file structure to work with
filesAndHashes = {}
print("Found "+str(len(filesToWatch))+" file(s) to watch:")
for filename in filesToWatch:
    print("Watching "+filename)

print(" ")
print("Use CTRL-C to exit this watcher ")
print(" ")

#create the archive folder if we need to
archiveFolderPath = os.path.join(".",archiveFolderName)
if not os.path.exists(archiveFolderPath):
    os.mkdir(archiveFolderPath)

sys.stdout.flush() #had to flush to make sure the console stays updated before it starts spinning on the files

#main thread
try:
    while True:
        for filename in filesToWatch:
            #create a blank hash if the hash list hasn't seen the file yet
            if(filename not in filesAndHashes):
                filesAndHashes[filename] = ""

            try:
                #open and read in the file
                file = open(filename, 'rb')
                filecontent = file.read()
                file.close();

                #hash the file contents
                h = md5()
                h.update(str(filecontent).encode("utf-8"))

                #if the latest hash doesn't match the last hash for the file
                #then archive the file with a timestamp in the file name
                if(h.hexdigest() != filesAndHashes[filename]):
                    filesAndHashes[filename] = h.hexdigest()

                    timestamp = time.localtime();
                    print("updated " + filename + " was backed up at " + str(time.strftime("%H:%M:%S",timestamp)))

                    newFileName = str(time.strftime("%Y%m%d_%H%M%S",timestamp)) + "-" + filename
                    newFilePath = os.path.join(archiveFolderPath,newFileName);
                    newFile = open(newFilePath, mode="wb")
                    newFile.write(filecontent)
                    newFile.close()
            #IO Errors should throw a soft warning but aren't cause for alarm; we may get it backed up next pass
            except IOError  as e:
                print("burped on "+filename+" with taste: "+str(e))

            sys.stdout.flush()
            time.sleep(secondsBetweenRefreshes);

except(KeyboardInterrupt): #CTRL-C to exit
    sys.exit(0);