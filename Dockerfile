FROM debian
ADD ./webtest .
EXPOSE 3000
CMD ./webtest
