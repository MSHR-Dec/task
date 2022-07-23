import type { NextPage } from 'next'
import Image from 'next/image'
import styles from '../styles/Home.module.css'
import {Button, Input, Spacer} from "@nextui-org/react"
import {UserContext} from "./_app";
import {useContext} from "react"
import { useRouter } from 'next/router'

/**
 * Note:
 *  Warning: Prop `id` did not match. Server: "react-aria-2" Client: "react-aria-4"
 *  See: https://issuemode.com/issues/nextui-org/nextui/53755914
 */
const Top: NextPage = () => {
    const router = useRouter()
    const { state, dispatch } = useContext(UserContext)

    // @ts-ignore
    const handleSubmit = async (event) => {
        event.preventDefault();

        const res = await fetch(`http://localhost:8080/signin`,
            {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(
                    {
                        "name": event.target.name.value,
                        "password": event.target.password.value,
                    }
                ),
                credentials: "same-origin",
            }
        )
        const body = await res.json()
        dispatch({
            type: "SET_USER",
            user: {id: body.id, name: event.target.name.value}
        })
        await router.push(`/home/${body.id}`)
    }

    return (
      <div className={styles.container}>
          <form onSubmit={handleSubmit}>
              <Spacer y={2} />
              <Input clearable id="name" label="name" placeholder="Name" />
              <Spacer y={2} />
              <Input.Password id="password" label="password" placeholder="Password" />
              <Spacer y={2} />
              <Button type="submit">Sign in !</Button>
              <Spacer y={2} />
          </form>

          <footer className={styles.footer}>
                <a
                    href="https://github.com/MSHR-Dec/task"
                    target="_blank"
                    rel="noopener noreferrer"
                >
                    Powered by{' '}
                    <span className={styles.logo}>
                        <Image src="/task.svg" alt="Task Logo" width={72} height={16} />
                    </span>
                </a>
          </footer>
      </div>
    )
}

export default Top
