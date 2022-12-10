import React from "react";
import {Container} from "semantic-ui-react";
import ToDoList from "./todo-list"
import './App.css';

function App() {
  return (
    <div>
      <Container>
        <ToDoList/>
      </Container>
    </div>
  );
}

export default App;
