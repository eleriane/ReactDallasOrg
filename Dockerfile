FROM scratch
ADD static /static
ADD ReactDallasOrg /
ENTRYPOINT ["/ReactDallasOrg"]