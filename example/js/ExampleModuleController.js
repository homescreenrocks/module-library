    angular.module('homescreenApp').controller('ExampleModuleController', ['$http', '$q', ExampleModuleController]);

    function ExampleModuleController($http, $q) {
        var vm = this;
        vm.test = 'text';
    }
