import { render } from 'react-dom'
import { ApolloClient, InMemoryCache, ApolloProvider } from '@apollo/client'
// import { persistCache, localStorageWrapper } from 'apollo3-cache-persist'
import App from './App'
import { ThemeProvider } from './contexts/theme'

import './index.css'

console.log(process.env.REACT_APP_GRAPHQL_URL)
const client = new ApolloClient({
  uri: process.env.REACT_APP_GRAPHQL_URL,
  cache: new InMemoryCache(),
})
render(
  <ApolloProvider client={client}>
    <ThemeProvider>
      <App />
    </ThemeProvider>
  </ApolloProvider>,
  document.getElementById('root')
)
