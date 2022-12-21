import {Component} from 'solid-js';
import './App.scss';
import {Link, useLocation, useRoutes} from "@solidjs/router";
import {routes} from "./routes";


const App: Component = () => {
  const location = useLocation()
  const Route = useRoutes(routes)
  return(
    <>
      <nav>
        <ul>
          <li >
            <Link href="/">
              Home
            </Link>
          </li>
          <li >
            <Link href="/about">
              About
            </Link>
          </li>
          <li >
            <Link href="/error" >
              Error
            </Link>
          </li>

          <li >
            <span>URL:</span>
            <input
              type="text"
              readOnly
              value={location.pathname}
            />
          </li>
        </ul>
      </nav>

      <main>
        <Route />
      </main>
    </>
  )
};

export default App;
