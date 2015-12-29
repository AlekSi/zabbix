#!/usr/bin/env ruby

# Common setup script for Vagrant and Travis CI.

ZABBIX_VERSION = ENV['ZABBIX_VERSION']
puts "Installing Zabbix #{ZABBIX_VERSION} ..."

def run(cmd)
  puts "# #{cmd}"
  res = system cmd
  if not res
    puts "#{res} #{$?}"
    exit $?.to_i
  end
end

run('wget -qO - http://repo.zabbix.com/zabbix-official-repo.key | apt-key add -')
run("add-apt-repository 'deb http://repo.zabbix.com/zabbix/#{ZABBIX_VERSION}/ubuntu/ trusty main non-free contrib'")
run('apt-get update -qq')
run('apt-get install -y postgresql')
run('apt-get install -y apache2 libapache2-mod-php5 php5-pgsql')
File.open('/etc/php5/apache2/php.ini', 'a') do |f|
  f.puts '[Date]'
  f.puts 'date.timezone = UTC'
end
run('apt-get install -y zabbix-server-pgsql zabbix-frontend-php')

conf = File.read('/etc/dbconfig-common/zabbix-server-pgsql.conf')
password = /dbc_dbpass='(\w+)'/.match(conf)[1]

File.open('/usr/share/zabbix/conf/zabbix.conf.php', 'w') do |f| f.puts <<-END
  <?php
  // Zabbix GUI configuration file
  global $DB;

  $DB['TYPE']     = 'POSTGRESQL';
  $DB['SERVER']   = 'localhost';
  $DB['PORT']     = '0';
  $DB['DATABASE'] = 'zabbix';
  $DB['USER']     = 'zabbix';
  $DB['PASSWORD'] = '#{password}';

  // SCHEMA is relevant only for IBM_DB2 database
  $DB['SCHEMA'] = '';

  $ZBX_SERVER      = 'localhost';
  $ZBX_SERVER_PORT = '10051';
  $ZBX_SERVER_NAME = '';

  $IMAGE_FORMAT_DEFAULT = IMAGE_FORMAT_PNG;
  ?>
  END
end
