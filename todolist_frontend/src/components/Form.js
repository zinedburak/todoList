import React from "react"
import nextId from "react-id-generator"

const Form = ({inputText, setInputText, todos, setTodos}) =>{
    const inputTextHandler = (e) => {
        setInputText(e.target.value)
    }
    const submitTodoHandler = async (e) => {
        const todoId = nextId();
        setTodos([...todos,{text: inputText, id: todoId}]);
        const response = await fetch('http://localhost:8000/api/add_todos',{
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                text: inputText,
                completed: 'false',
            })
        });
        const content = await response.json();
        
    }

    return (
        <form>
            <input value={inputText} onChange={inputTextHandler} type="text" className="todo-input" placeholder="Add Todo" />
            <button onClick={submitTodoHandler} className="todo-button" type="submit" data-testid="add-button">
                <i className="fas fa-plus-square"></i>
            </button>
        </form>
    )
}

export default Form;