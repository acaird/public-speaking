%title: Security Keynote
%author: Kris Nova [Twitter @krisnova]
%date: 1992-01-01

-> Cloud Native Security <-
=========

-> # Kris Nova <-

-> github.com/kris-nova/public-speaking <-

---

-> # About Me <-

_Author:_
	Cloud Native Infrastructure
	nivenly.com
	Blogs
	CNCF
	Github

_Contibutor:_
	Kubernetes
		kops, kubeadm, kubicorn, kubernetes, kubectl, cluster-api
	Go standard library
	Linux, FreeBSD
	Terraform
	Falco
	Sysdig
	Wireshark

---

-> # Problem <-

_How do I monitor my system in such a way that I cannot break it?_
_How can I ensure that nothing can bypass our tooling?_
	Audit from the network?
	Audit from the kernel?
	Making any data meaningful?

_Smaller problem:_
	How can I make a presentation from the command line...
	...that references man pages and C...	*FUN*

---

-> # Agenda <-

_Review:_
	*Understanding Containers*
	1) cgroups(7)
	2) namespaces(7)
	3) clone(2)
	*Securing Containers*
	1) Kernel Modules vs eBPF
	2) Falco

_Demonstrate:_
	Writing container code in C
	Auditing with Linux
	Auditing in a Docker Container
	Auditing wth Kubernetes (*We are literally going to hack into a cluster*)

---

-> # cgroups(7) <-
	Control groups, usually referred to as cgroups, are a Linux kernel feature which
	allow processes to be organized into hierarchical groups whose usage of various
	types of resources can then be *limited and monitored*.
	
	 $ cd /sys/fs/cgroup
	 $ ps aux | grep <pattern>
	 $ echo <pid> > <subsystem>/<limit>

	 dr-xr-xr-x 7 root root  0 Oct 27 15:37 blkio
	 dr-xr-xr-x 7 root root  0 Oct 27 15:37 cpu,cpuacct
	 dr-xr-xr-x 3 root root  0 Oct 27 15:37 cpuset
	 dr-xr-xr-x 7 root root  0 Oct 27 15:37 devices
	 dr-xr-xr-x 3 root root  0 Oct 27 15:37 freezer
	 dr-xr-xr-x 3 root root  0 Oct 27 15:37 hugetlb
	 dr-xr-xr-x 7 root root  0 Oct 27 15:37 memory
	 dr-xr-xr-x 3 root root  0 Oct 27 15:37 net_cls,net_prio
	 dr-xr-xr-x 3 root root  0 Oct 27 15:37 perf_event
	 dr-xr-xr-x 7 root root  0 Oct 27 15:37 pids
	 dr-xr-xr-x 2 root root  0 Oct 27 15:37 rdma
	 dr-xr-xr-x 7 root root  0 Oct 27 15:37 systemd
	 dr-xr-xr-x 6 root root  0 Oct 27 15:37 unified

---

-> # namespaces(7)
	A namespace wraps a global system resource in an abstraction that makes it *appear*
	to the processes within the namespace that they have their own isolated instance
	of the global resource

	Namespace Flag            Page                  Isolates
	Cgroup    CLONE_NEWCGROUP cgroup_namespaces(7)  Cgroup root directory
	IPC       CLONE_NEWIPC    ipc_namespaces(7)     System V IPC, POSIX message queues

	Network   CLONE_NEWNET    network_namespaces(7) Network devices, stacks, ports, etc
	
	Mount     CLONE_NEWNS     mount_namespaces(7)   Mount points
	PID       CLONE_NEWPID    pid_namespaces(7)     Process IDs
	User      CLONE_NEWUSER   user_namespaces(7)    User and group IDs
	UTS       CLONE_NEWUTS    uts_namespaces(7)     Hostname and NIS

---

-> # clone(2)
	clone creates a new child process, and can share parts of it's
	execution context with this process.

	container:
		creating a process with clone(2) and managing with a cgroup

---

-> # clone(20) example <-
	// /* Same Pid, Same Disk */
	int pid = clone(fn, pchild_stack + (1024 * 1024), SIGCHLD, NULL);
	//
	// /* Different Pid, Same Disk */
	//int pid = clone(fn, pchild_stack + (1024 * 1024),
	      CLONE_NEWPID | SIGCHLD, NULL);
	//
	// /* Different Pid, Different Disk */
	//int pid = clone(fn, pchild_stack + (1024 * 1024),
	      CLONE_NEWPID | CLONE_NEWNET | CLONE_NEWNS | SIGCHLD, NULL);
	      
---

-> # Review <-
	What have we learned?
	
	1) Containers are processes running with namespaces(7) and cgroups(7)
	2) Containers can be built in a small amount of code
	3) Kubernetes runs containers
	4) Container runtimes *and* Kubernetes abstracts scary things

-> **DEMO TIME!** <-

-> EOF; <-