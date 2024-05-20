FROM grafana/grafana
ENV GF_AUTH_DISABLE_LOGIN_FORM "true"
ENV GF_AUTH_ANONYMOUS_ENABLED "true"
ENV GF_AUTH_ANONYMOUS_ORG_ROLE "Admin"
#ADD .grafana/provisioning /etc/grafana/provisioning
#ADD .grafana/config.ini /etc/grafana/config.ini
#ADD .grafana/dashboards /var/lib/grafana/dashboards

FROM golang:1.22.3-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./cmd/api
EXPOSE 8081
CMD ["./main"]