const CheckLogin = () => {
    return localStorage.getItem("accessToken")

}
const Logout = () => {
    localStorage.removeItem("accessToken");
}

export {
    CheckLogin,
    Logout,
}


