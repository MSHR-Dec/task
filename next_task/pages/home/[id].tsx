import {useContext} from "react"
import {UserContext} from "../_app"
import {parseCookies} from "nookies";
import {Card, Spacer, Text} from "@nextui-org/react";

type Task = {
    id: number
    name: string
    status: string
    startAt: string
    endAt: string
    createdAt: string
    modifiedAt: string
}

type HomeProps = {
    tasks: Task[]
}

const Home = (props: HomeProps) => {
    const {state} = useContext(UserContext)
    const cookie = parseCookies()
    console.log(state)
    console.log(cookie)
    return (
        <>
            {props.tasks.map((task) => (
                <>
                    <Card>
                        <Card.Body>
                            <Text>{task.name}</Text>
                        </Card.Body>
                        <Card.Footer>
                            <Text>{task.status}</Text>
                        </Card.Footer>
                    </Card>
                    <Spacer y={1}/>
                </>
            ))}
        </>
    )
}

export const getServerSideProps = async (context: any): Promise<{ props: HomeProps }> => {
    const cookie = parseCookies(context)
    const res = await fetch(`http://backend:8080/users/${context.query.id}/tasks`, {
        credentials: "include",
        headers: {cookie: `task=${cookie.task}`}
    })
    const body = await res.json()
    return {props: {tasks: body.tasks}}
}

export default Home
