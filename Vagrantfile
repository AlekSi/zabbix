# Vagrant box for local testing.

Vagrant.configure(2) do |config|
  config.vm.box = "ubuntu/precise64"
  config.vm.box_version = "20151217.0.0"

  config.vm.network "forwarded_port", guest: 80, host: 8080

  config.vm.provision "shell", inline: <<-SHELL
    cd /vagrant
    export DEBIAN_FRONTEND=noninteractive
    export ZABBIX_VERSION=2.4
    ./setup.rb
  SHELL
end
