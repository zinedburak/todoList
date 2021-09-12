import React, { useState, useEffect } from 'react'
import './App.css';
import Form from './components/Form';
import TodoList from './components/TodoList'

function App() {
  const [inputText, setInputText] = useState("");
  const [todos, setTodos] = useState([]);

  useEffect(() => {
    
    const getTodos =   async () =>{
           const response = await fetch('http://localhost:8000/api/get_todos',{            
               headers: {'Content-Type': 'application/json'},
               credentials: 'include',
           });
           const content = await response.json();
           console.log(content)
           setTodos(content)
           
       }
       getTodos();
 },[] )


  return (
    <div className="App">
      <header>
        <h1>Burak's Todo List </h1>
      </header>
      <Form inputText={inputText} setInputText={setInputText} todos={todos} setTodos={setTodos} />
      <TodoList todos={todos}/>
    </div>
    
  );
}

export default App;
