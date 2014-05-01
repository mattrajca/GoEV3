GoEV3
=====

Introduction
------------

In 2013, LEGO introduced its third generation Mindstorms robotics set, the [EV3](http://en.wikipedia.org/wiki/Lego_Mindstorms_EV3). Unlike its predecessors, the EV3 runs Linux, giving hackers and hobbyists the opportunity create robots more capable than ever before.

The [ev3dev project]() maintains open source, hacker-friendly releases of EV3's operating system. Distributions include built-in [ssh](http://en.wikipedia.org/wiki/Secure_Shell) support and custom drivers for EV3's hardware. In fact, a simple file system-based interface can be used to interact with EV3's motors, sensors, buttons, and LEDs. Directories under `/sys/class` represent various device classes, and setting attributes is as simple as writing to files.

For example, executing the following shell commands will run a motor at 50% speed:

	echo  50 > /sys/class/tacho-motor/outA:motor:tacho/speed_setpoint
	echo   1 > /sys/class/tacho-motor/outA:motor:tacho/run

This enables third party developers to write EV3 bindings for any programming language/framework that has a file system IO API.

GoEV3 provides EV3 bindings for [Google Go](http://golang.org), allowing developers to take advantage of Go's modern syntax and extensive standard library while programming Mindstorms robots.

Getting Started
---------------

### ev3dev

First, we need to install ev3dev onto a Micro SD card (by using an SD card, we can keep EV3's built-in software intact). Instructions for the installation process can be found on [ev3dev's wiki](https://github.com/mindboards/ev3dev/wiki/Getting-started-v2). When you're done, reboot your EV3 and make sure you can ssh into it from your computer.

### Google Go

Next, we need to install an ARMv5 build of Google Go. Fortunately for us, Go developer Dave Cheney released [builds](http://dave.cheney.net/unofficial-arm-tarballs) of Go for various ARM architectures. On your computer, download the ARMv5 package. It works out of the box on the EV3! Once the download completes, transfer it to the EV3 over ssh using [scp](http://en.wikipedia.org/wiki/Secure_copy):

	scp /path/to/go1.2.linux-arm~armv5-1.tar.gz root@192.168.3.2:~/go.tar.gz

Be sure to replace `192.168.3.2` with your EV3's IP address. Now we can ssh into the EV3 and extract the archive to its final destination:

	cd /usr/local
	tar -xf ~/go.tar.gz

Extraction may take a few minutes. Lastly, we'll add Go's `bin` directory to our shell's path:

	echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
	source ~/.bashrc

You should now be able to invoke the `go` tool like so:

	root@ev3dev:~# go version
	go version go1.2 linux/arm

### GoEV3

Now that we have Google Go up and running, we need to install GoEV3. First, let's set up our Go workspace:

	cd ~
	mkdir gocode
	echo "export GOPATH=\$HOME/gocode" >> ~/.bashrc
	source ~/.bashrc

We can easily obtain GoEV3 from its GitHub repository. Be sure to have internet connection sharing enabled prior to running the following command:

	mkdir -p gocode/src/github.com/mattrajca
	cd gocode/src/github.com/mattrajca
	wget -O GoEV3.tar.gz --no-check-certificate https://github.com/mattrajca/GoEV3/archive/master.tar.gz
	tar -xf GoEV3.tar.gz
	mv GoEV3-master GoEV3
	rm GoEV3.tar.gz
	cd ~

We should now be able to run the sample program included with GoEV3.

	go build github.com/mattrajca/GoEV3
	gocode/bin/GoEV3

Choose option `6. Motors`, plug in a motor to output port A, and watch it turn! Feel free to explore the other modes.

Your First Program
------------------

Documentation
-------------

Complete documentation for GoEV3 can be found on [godoc](https://godoc.org/github.com/mattrajca/GoEV3).

Contributing
------------

GoEV3 is still in its early stages and subject to API changes as the ev3dev project evolves. Filing issues and submitting pull requests are the two best ways to get involved. Documentation improvements, new APIs, example programs, and bug fixes are all welcome.
