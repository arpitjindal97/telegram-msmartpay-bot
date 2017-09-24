FROM openjdk:latest

RUN apt-get update \
	&& apt-get install -y --no-install-recommends iceweasel xvfb xauth

RUN mkdir /arpit
COPY . /arpit/

RUN wget -O geckodriver.tar.gz \
	$(curl -s https://api.github.com/repos/mozilla/geckodriver/releases/6998290 | grep browser_download_url | grep linux64 | cut -d '"' -f 4)

RUN tar -xzvf geckodriver.tar.gz
RUN cp geckodriver /arpit/ && rm geckodriver.tar.gz

RUN wget -O selenium-server-standalone.jar https://goo.gl/s4o9Vx
RUN cp selenium-server-standalone.jar /arpit/

RUN chmod +x /arpit/entrypoint.sh
ENTRYPOINT ["/arpit/entrypoint.sh"]
