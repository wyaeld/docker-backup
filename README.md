# Backup and Restore Docker Volume Containers

For general documentation on how to use volume containers, see:
http://docs.docker.io/en/latest/use/working_with_volumes/#creating-and-mounting-a-data-volume-container


Let say you have a container named `mysql-data` to keep `/var/lib/mysql`. You start up your mysql server by running:

    $ docker run --volumes-from=mysql-data --name mysql-server ...


Backup that data container:

    $ docker-backup store mysql-server mysql-server-backup.tar

Restore it on a new system:

    $ docker-backup restore mysql-server mysql-server-backup.tar

