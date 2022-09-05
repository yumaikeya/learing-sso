import { memo, useState, useEffect } from "react"
import { Auth } from "@aws-amplify/auth"
import { useNavigate } from 'react-router-dom'


const Top = memo(() => {
  const navigate = useNavigate()
  const [user, setUser] = useState({})

  const handleClick = () => {
    (async () => {
      try {
        await Auth.signOut()
        await Auth.currentSession()
      } catch (error) {
        navigate("/signin")
      }
    })()
  }

  useEffect(() => {
    (async () => {
      try {
        setUser(await Auth.currentAuthenticatedUser())
      } catch (error) {
        navigate("/signin")
      }
    })()
  }, [navigate])

  return (
    <>
      <h2 style={{marginTop: 0}}>Hello {user.username}</h2>
      <button onClick={handleClick}>SignOut</button>
    </>
  )
})

export default Top