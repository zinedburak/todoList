import React from "react"
import {render,fireEvent, queryByPlaceholderText} from '@testing-library/react'
import Form from "../Form"

it("renders correctly", () => {

    const {queryByTestId, queryByPlaceHolderText} = render(<Form />);
    expect(queryByTestId("add-button")).toBeTruthy();
    expect(queryByPlaceHolderText("Add Todo")).toBeTruthy();

});


// describe("Input Value", () => {
//     it("updates on change", () => {
//         const {queryByPlaceHolderText} = render([Form])

//         const todoInput = queryByPlaceHolderText("Add Todo")

//         fireEvent.change(todoInput, {target: {value: "test"}})

//         expect(todoInput.value).toBe("test")
//     })
// })