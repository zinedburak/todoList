import puppeteer from 'puppeteer'
import "@testing-library/jest-dom/extend-expect"



    it("Passes Acceptance Test", async () => {
        const browser = await puppeteer.launch();
        const page = await browser.newPage();
        await page.goto("http://localhost:3000/");
        
        await page.click('#input_text');
        await page.type('#input_text', 'test1');
        await page.click('#add-button') 
        await page.waitForResponse(res => res.url().endsWith('add_todos'))
        await page.waitForResponse(res => res.url().endsWith('get_todos'))
        const todoItem = await page.$('#todo-item')
        const todoItemValue = await todoItem.evaluate((el) => el.textContent)

        expect(todoItemValue).toContain("test1")

        browser.close();
        
    })


