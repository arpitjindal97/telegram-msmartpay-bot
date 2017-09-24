FROM openjdk:latest

RUN apt-get update \
	#&& export TERM="xterm-256color" \
	&& apt-get install -y --no-install-recommends iceweasel xvfb xauth

RUN mkdir /arpit
COPY . /arpit/

RUN wget https://github.com/mozilla/geckodriver/releases/download/v0.18.0/geckodriver-v0.18.0-linux64.tar.gz

RUN tar -xzvf geckodriver-*-linux64.tar.gz
RUN cp geckodriver /arpit/ && rm geckodriver-*-linux64.tar.gz

RUN wget -O selenium-server-standalone.jar https://goo.gl/s4o9Vx
RUN cp selenium-server-standalone.jar /arpit/

RUN chmod +x /arpit/entrypoint.sh
ENTRYPOINT ["/arpit/entrypoint.sh"]
