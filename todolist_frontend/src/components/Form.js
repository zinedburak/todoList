import React from "react"
import nextId from "react-id-generator"

const Form = ({inputText, setInputText, todos, setTodos}) =>{
    const inputTextHandler = (e) => {
        setInputText(e.target.value);
        console.log("Burak")
        console.log(inputText)
    }
    const submitTodoHandler = (e) => {
        e.preventDefault();
        console.log("inputText")
        
        const todoId = nextId();
        setTodos([...todos,{text: inputText, id: todoId}]);
        setInputText("5")
        
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