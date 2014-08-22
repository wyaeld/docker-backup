# Backup and Restore Data Volume Containers

Let say you have a Data Volume Container named `mysql-data` to keep `/var/lib/mysql`. You start up your mysql server container by running:

    $ docker run --volumes-from=mysql-data --name mysql-server ...

This tool allows you to backup and restore the Data Volume Container `mysql-data` easily.

## Backup
There are two ways to backup a Data Volume Container:

1. Via a Data Volume Container Id or Name directly

	```
	$ docker-backup store mysql-data-backup.tar mysql-data
	```

2. Via a Container Id or Name
	In this case the tool will try to find the Data Volume Container for the given container.

	```
	$ docker-backup store mysql-data-backup.tar mysql-server
	```

Contents of the backup tarball:

 * the Data Volume Container's json
 * all volumes found in the Data Volume Container

## Restore
After that, this tool can be used to restore a Data Volume Container from that tarball
either on the same or on another host.

    $ docker-backup restore mysql-data-backup.tar

For general documentation on how to use Data Volume Containers, see:
http://docs.docker.io/en/latest/use/working_with_volumes/#creating-and-mounting-a-data-volume-container

For a more complete backup strategy built on top of docker-backup, look at https://github.com/discordianfish/docker-lloyd
