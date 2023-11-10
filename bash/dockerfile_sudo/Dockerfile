FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y golang-go

RUN git clone https://github.com/hybridgroup/gocv.git

# remove sudo
RUN sed -i 's/sudo //g' gocv/Makefile

RUN cd gocv && \
    make install && \
    make clean
