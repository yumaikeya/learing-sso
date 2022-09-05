import { useState } from "react"
import { Auth } from "@aws-amplify/auth"
import Singin from "./components/Signin"
import Top from "./components/Top"
import { BrowserRouter, Route, Routes, Navigate, Outlet } from "react-router-dom"


Auth.configure({
  region: process.env.REACT_APP_REGION,
  userPoolId: process.env.REACT_APP_USER_POOL_ID, // INSPECTION-Dev
  userPoolWebClientId: process.env.REACT_APP_USER_POOL_CLIENT_ID,
  identityPoolId: process.env.REACT_APP_ID_POOL_ID,
  mandatorySignIn: true,
  authenticationFlowType: "USER_SRP_AUTH", // USER_PASSWORD_AUTH || USER_SRP_AUTH
  cookieStorage: {
    path: "/",
    expires: 1,
    sameSite: "lax",
    secure: false,
    domain: "localhost"
  }
})

const App = () => {
  const [isAuthenticated, setAuthenticated] = useState(true)

  const PrivateRoute = () => {
    (async () => {
      try {
        await Auth.currentAuthenticatedUser()
        setAuthenticated(true)
      } catch (error) {
        setAuthenticated(false)
      }
    })()
  
    return isAuthenticated
      ? <Outlet />
      : <Navigate to="/signin" replace />
  }

  const ValidateRoute = () => {
    (async () => {
      try {
        await Auth.currentAuthenticatedUser()
        setAuthenticated(true)
      } catch (error) {
        setAuthenticated(false)
      }
    })()
  
    return isAuthenticated
      ? <Navigate to="/" replace />
      : <Outlet />
  }

  return (
    <BrowserRouter>
      <Routes>
        <Route element={<PrivateRoute />}>
          <Route path="/" element={<Top/>} />
        </Route>
        <Route element={<ValidateRoute />}>
          <Route path="/signin" element={<Singin/>} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App
