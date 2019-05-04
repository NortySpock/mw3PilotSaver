import sys,os, time,hashlib
from hashlib import md5

initalListOfFiles = []
filesToWatch = dict()
userRequestedFile = ""
targetFileExtension = ".txt"
secondsBetweenRefreshes = 3
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
    h = md5()
    h.update(str(filename).encode("utf-8")) #TODO use file contents    
    filesAndHashes[filename] = h.hexdigest()
    print("Watching "+filename+" with MD5 hash of " + h.hexdigest())

print(" ")

#create the archive folder if we need to
archiveFolderPath = os.path.join(".",archiveFolderName)
if not os.path.exists(archiveFolderPath):
    os.mkdir(archiveFolderPath)

#testing the time
print("local time: " + str(time.strftime("%H:%M:%S")))

print("filename: " + str(time.strftime("%Y%m%d_%H%M%S")))

#main thread
# try:
    # while True:
        # time.sleep(secondsBetweenRefreshes) 
        # //print("running...")
# except KeyboardInterrupt:
    # sys.exit(0);