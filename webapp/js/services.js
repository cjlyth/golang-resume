'use strict';
/* Services */
angular.module('resumeApp.services', ['ngResource'])
.value('companies', ["booz", "army", "publicrelay" ])
.factory('resumeData', ['$resource', function($resource) {
	var CompanyExperience = $resource('js/:company.json'
													, { company: '@id'}
													, {
														get: {
															method: 'GET',
															cache: false
														}
													});
	function getExperience(company){
		if (angular.isString(company)) {
			return getExperience({company: company});
		} else if (angular.isArray(company)) {
			return company.map(getExperience);
		} else if (angular.isObject(company)) {
			return CompanyExperience.get(company);
		} else {
			throw "Invalid argument to getExperience: " + typeof(company);
		}
	}
	return { getExperience: getExperience };
}])
;
