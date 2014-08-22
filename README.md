# Backup and Restore Docker Volume Containers

This tool gets the [volume container](http://docs.docker.io/en/latest/use/working_with_volumes/#creating-and-mounting-a-data-volume-container) of a given container and writes the volume container's json and all volumes found to a tarball.
After that, this tool can be used to restore a volume container from that tarball.

Let say you have a container named `mysql-data` to keep `/var/lib/mysql`. You start up your mysql server by running:

    $ docker run --volumes-from=mysql-data --name mysql-server ...


Backup that data container:

    $ docker-backup store mysql-server-backup.tar mysql-server

Restore it on a new system:

    $ docker-backup restore mysql-server-backup.tar


For general documentation on how to use volume containers, see:
http://docs.docker.io/en/latest/use/working_with_volumes/#creating-and-mounting-a-data-volume-container

For a more complete backup strategy built on top of docker-backup, look at https://github.com/discordianfish/docker-lloyd
