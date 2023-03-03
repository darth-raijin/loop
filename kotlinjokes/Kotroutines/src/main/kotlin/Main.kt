import kotlinx.coroutines.runBlocking
import okhttp3.*
import java.io.IOException


fun main(args: Array<String>) {
    var client = OkHttpClient()
    val routineJob = runBlocking {
        sendGetRequest("", client)
    }

    

    routineJob.join()
}


suspend fun sendGetRequest(url: String, client: OkHttpClient): Int {
    val request: Request = Request.Builder()
        .url(url)
        .build()

    var foo: Any

    val response = client.newCall(request).enqueue(object: Callback {
        override fun onFailure(call: Call, e: IOException) {
            TODO("Not yet implemented")
        }

        override fun onResponse(call: Call, response: Response) {
            response.code
        }

    })

    return 2
}
