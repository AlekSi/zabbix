# Vagrant box for local testing.

Vagrant.configure(2) do |config|
  config.vm.box = "ubuntu/trusty64"
  config.vm.box_version = "20160217.1.0"

  config.vm.hostname = "vagrant"

  config.vm.network "forwarded_port", guest: 80, host: 8080

  config.vm.provision "shell", inline: <<-SHELL
    cd /vagrant
    export DEBIAN_FRONTEND=noninteractive
    export ZABBIX_VERSION=2.4
    ./setup.rb
  SHELL
end
