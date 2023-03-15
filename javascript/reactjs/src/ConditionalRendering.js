const root = document.querySelector("#root"); //

// Component App
function App() {
    // State "login" with value "false"
    const [isLogin, setLogin] = React.useState(false);
    console.log(isLogin);

    // Conditional Rendering
    // Jika isLogin == true, maka dieksekusi
    // jika isLogin == false, maka lewati statement ini
    if (isLogin) {
        return (
            <>
                <h1 style={{ color: "green" }}>Sudah Login</h1>
                <button
                    onClick={function () {
                        setLogin(false);
                    }}>
                    Logout
                </button>
            </>
        );
    }

    return (
        <>
            <h1>Belum Login</h1>
            {/* Bisa pakai In Line Condition */}
            {/* {isLogin ? <h1>Silahkan Melanjutkan</h1> : null} */}
            <button
                onClick={function () {
                    setLogin(true);
                }}>
                Login
            </button>
        </>
    );
}

ReactDOM.render(<App />, root);
