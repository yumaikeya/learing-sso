import { memo } from "react"
import { Auth } from '@aws-amplify/auth'
import { useNavigate } from 'react-router-dom'
// import { isEmpty } from '../utils/comparer'

const Singin = memo(() => {
  const navigate = useNavigate()

  const handleClick = () => {
    (async () => {
      try {
        await Auth.signIn(
          process.env.REACT_APP_USER_NAME,
          process.env.REACT_APP_USER_PASSWORD
        )

        await Auth.currentSession()
        console.log("SignIn done!")
        navigate("/")
      } catch (error) {
        console.log('SignIn err: ', error)
        navigate("/signin")
      }
    })()  
  }

  return (
    <>
      <button onClick={handleClick}>SignIn</button>
      <h2>↑↑ You need to signin</h2>
    </>
  )
})

export default Singin