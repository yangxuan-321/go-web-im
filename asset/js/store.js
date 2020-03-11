function LocalStore(){

}

LocalStore.setTokenInfo = function (data) {
    localStorage.token = data.token;
    localStorage.userId = data.userId;
}

LocalStore.getTokenInfo = function () {
    return {userId: localStorage.getItem('userId'), token: localStorage.getItem('token')};
}
