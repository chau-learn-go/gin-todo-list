async function displayJob() {
    const jobsResponse = await loadJobs()
    const jobs = jobsResponse.items

    const oldTable = document.querySelector("table")
    if (oldTable) {
        oldTable.remove()
    }

    const table = document.createElement("table")
    const thead = table.createTHead()
    const headerRow = thead.insertRow()

    const headers = ["Index", "Job Name", "Status", "Actions"]

    headers.forEach(headerText => {
        const th = document.createElement("th")
        th.textContent = headerText
        headerRow.appendChild(th)
    })

    const tbody = table.createTBody()

    jobs.forEach((job, index) => {
        const row = tbody.insertRow()

        const cellIndex = row.insertCell()
        cellIndex.textContent = index + 1

        const cellName = row.insertCell()
        cellName.textContent = job.name

        const cellStatus = row.insertCell()
        cellStatus.textContent = job.status

        const cellActions = row.insertCell()

        if (job.status !== "doing" && job.status !== "done") {
            const btnDoing = document.createElement("button")
            btnDoing.textContent = "Doing"
            cellActions.appendChild(btnDoing)
        }

        if (job.status !== "done") {
            const btnDone = document.createElement("button")
            btnDone.textContent = "Done"
            cellActions.appendChild(btnDone)
        }

        const btnDelete = document.createElement("button")
        btnDelete.textContent = "Delete"
        cellActions.appendChild(btnDelete)
    })

    const app = document.getElementById("app")
    app.append(table)
}

function start() {
    displayJob()
}

start()