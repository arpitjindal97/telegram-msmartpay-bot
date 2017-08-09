
set :docker_repo, "arpitjindal1997/telegram-msmartpay"

task :deploy do
	on "root@139.59.71.230" do
		execute "export docker_cont_id=$(docker ps -a | grep #{fetch(:docker_repo)} | cut -d ' ' -f 1 ) && docker stop $docker_cont_id && docker rm $docker_cont_id"
		execute "docker rmi #{fetch(:docker_repo)}"
		execute "docker pull #{fetch(:docker_repo)} > /dev/null"
		execute "docker run -d #{fecth(:docker_repo)}"
	end
end
