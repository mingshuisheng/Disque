import {Component} from 'solid-js';
import './App.scss';
import {useLocation, useRoutes} from "@solidjs/router";
import {routes} from "./routes";


const App: Component = () => {
  const Route = useRoutes(routes)
  return (
    <main class={"root"}>
      <Route/>
    </main>
  )
};

export default App;
