# static web server

A static web server: a webserver that serves static files.  

## Conversion into a Production Server

steps to add from the simple stadard library server:  

### base server
web server that serves static files from the root directory specified in the Cl.  

./staticServer /port=portStr /root=rootDir /dbg  

### planned additions

1. logging
2. tls
	a. certificates stored in files
	b. in-memory certificartes
3. file transfer
	a. read file followed by write to socket
	b. io.Copy
	c. io.Copy with buffer pulled from ring
	d. giouring
4. worker pool vs limits

### performance measurements and comparison


### references
to come
