async function loadJobs() {
    try {
        const response = await fetch("/api/job")
        return await response.json()
    } catch (error) {
        console.error("Error loading jobs:", error)
    }
}

async function addJob(jobName) {
    const response = await fetch("/api/job", {
        method: "POST",
        body: JSON.stringify({
            name: jobName
        })
    })
    return response.ok
}

async function doJob(jobName) {
    const response = await fetch(`/api/job/${jobName}/doing`, {
        method: "PUT"
    })
    return response.ok
}

async function doneJob(jobName) {
    const response = await fetch(`/api/job/${jobName}/done`, {
        method: "PUT"
    })
    return response.ok
}

async function deleteJob(jobName) {
    const response = await fetch(`/api/job/${jobName}`, {
        method: "DELETE"
    })
    return response.ok
}