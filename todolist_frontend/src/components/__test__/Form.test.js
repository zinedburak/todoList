import React, { useState } from "react"
import {render,fireEvent, queryByPlaceholderText} from '@testing-library/react'
import Form from "../Form"
import "@testing-library/jest-dom/extend-expect"



it("renders correctly", () => {
    
    const component = render(<Form/>)
    const button = component.getByTestId("add-button");
    const input_text = component.getByPlaceholderText("Add Todo");
    expect(button).toBeTruthy;
    expect(input_text).toBeTruthy;
    

});


describe("Input Value", () => {
    it("updates on change", () => {
        const setInputText = jest.fn();
        const component = render(<Form setInputText={setInputText}/>)
        const todoInput = component.getByPlaceholderText("Add Todo")

        fireEvent.change(todoInput, {target: {value: "test"}})

        expect(todoInput.value).toBe("test")
        expect(setInputText).toHaveBeenCalled()
    })
})

describe("Add Button" , () => {    
    it(" The function Triggers " , () => {
        const setTodos = jest.fn();
        const setInputText = jest.fn();
        const todos = new Array();
        
        const component = render(<Form setTodos={setTodos} todos={todos} setInputText={setInputText}/>)
        const todoInput = component.getByPlaceholderText("Add Todo")
        const addButton = component.getByTestId("add-button")

        fireEvent.change(todoInput, {target: {value: "test"}})
        fireEvent.click(addButton)

        expect(todoInput.value).toBe("test")
        expect(setInputText).toHaveBeenCalled()
        //expect(setTodos).toHaveBeenCalled()
    })

})