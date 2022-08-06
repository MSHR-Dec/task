import {useContext} from "react"
import {UserContext} from "../_app"
import {parseCookies} from "nookies";
import {Card, Grid, Spacer, Text} from "@nextui-org/react";

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
    return (
        <>
            {props.tasks.map((task) => (
                <>
                    <Spacer y={2}/>
                    <Card>
                        <Card.Header>
                            <Grid.Container gap={2} justify="center">
                                <Grid xs={4}>
                                    <Text>{task.startAt}</Text>
                                </Grid>
                                <Grid xs={2}>
                                    <Text>~</Text>
                                </Grid>
                                <Grid xs={4}>
                                    <Text>{task.endAt}</Text>
                                </Grid>
                            </Grid.Container>
                        </Card.Header>
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
    const tasks: Task[] = body.tasks.map((task: any) => {
        return {
            id: task.id,
            name: task.name,
            status: task.status,
            startAt: task.start_at,
            endAt: task.end_at,
            createdAt: task.created_at,
            modifiedAt: task.modified_at
        }
    })
    return {props: {tasks: tasks}}
}

export default Home
