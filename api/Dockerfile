FROM mcr.microsoft.com/dotnet/core/sdk:3.1

COPY . ./app
RUN dotnet pubish

FROM mcr.microsoft.com/dotnet/core/runtime:3.1

EXPOSE 80
ENTRYPOINT ./publish/api.exe