# archive
A little project to manage files


## Steps
- [x] Create a hash of a file
- [ ] Check if the hash already exists, if yes: skip to next file | no: follow the next step
    - [ ] SQL stuff
- [ ] Get file [metadata](#metadata)
- [ ] Save data in database
- [ ] Move file to destination folder 
- [ ] Go to the next file


## Metadata
- name
- size
- parent folder
- current folder
- hash file
- tags
    - Extract all words that is possible from a doc file
    - Compare those word with other files with same type, location and similar file name
    - Those words that comes common bettwen the files become tags



## Directives
- all files will be put in the parent folder based on your type Documents, Pictures, Videos, Code, Apps;
- the files keep child folder: Document/bills, Pictures/tripp, Apps/linux, etc;
- all images go to the Pictures folder, unless it comes from the Document parent folder
- all documents files (doc, pdf, csv, txt, etc) go to the Document folder with the subfolder eg. bill, study, etc
