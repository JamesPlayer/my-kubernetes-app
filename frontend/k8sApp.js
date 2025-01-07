function jsonToStr(json) {
    return JSON.stringify(json, null, 4);
}

angular.module('myK8sApp', [])
    .controller('PingController', function ($scope, $http) {
        const ping = this;
        ping.response = jsonToStr({});

        function onSuccess(response) {
            ping.response = jsonToStr(response.data);
        }

        function onFailure(response) {
            ping.response = jsonToStr({
                error: {
                    status: response.status,
                    statusText: response.statusText,
                    xhrStatus: response.xhrStatus,
                    data: response.data,
                } 
            });
        }

        ping.makeRequest = function () {
            ping.response = jsonToStr({});
            $http.get("/ping").then(onSuccess, onFailure);
        };
    });