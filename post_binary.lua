wrk.method = "POST"
wrk.headers["Content-Type"] = "multipart/form-data; boundary=--------------------------726156676786942976125223"
wrk.headers["Content-Disposition"] = "form-data; name=document; filename=1000.csv"

file = io.open("./csv/1000.csv", "rb")
wrk.body = file:read("*a")
