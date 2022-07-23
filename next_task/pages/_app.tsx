import '../styles/globals.css'
import type { AppProps } from 'next/app'
import {NextUIProvider} from '@nextui-org/react'
import {createContext, useReducer} from "react"

export type User = {
    id: number
    name: string
}

export type UserAction = {
    type: 'SET_USER'
    user: User
} | {
    type: 'REMOVE_USER'
}

export const userReducer = (state: User, action: UserAction) => {
    switch (action.type) {
        case 'SET_USER':
            return action.user
        case 'REMOVE_USER':
            return {id: 0, name: ""}
        default:
            return state
    }
}

type UserContext = {
    state: User;
    dispatch: any;
}

export const UserContext = createContext<UserContext>({
    state: {id: 0, name: ""},
    dispatch: null
})

function MyApp({ Component, pageProps }: AppProps) {
    const user = { id: 0, name: "" }
    const [state, dispatch] = useReducer(userReducer, user)

    return (
        <NextUIProvider>
            <UserContext.Provider value={{state, dispatch}}>
                <Component {...pageProps} />
            </UserContext.Provider>
        </NextUIProvider>
    )
}

export default MyApp
