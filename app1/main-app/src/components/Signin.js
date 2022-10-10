import { memo } from "react"
import { Auth } from '@aws-amplify/auth'
import { useNavigate } from 'react-router-dom'
// import { isEmpty } from '../utils/comparer'

const Singin = memo(() => {
  const navigate = useNavigate()

  const handleClick = () => {
    (async () => {
      try {
        await Auth.signIn({
          "username": process.env.REACT_APP_USER_NAME,
          "password": process.env.REACT_APP_USER_PASSWORD,
          "validationData": {
            "tenant": "INSPECTION"
          }
        })

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