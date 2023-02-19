package handler

/*
# DEFAULT_TIMEOUT = 60 seconds
# MAX_RETRIES = 3
# retries = 0
#
# DO
#     TRY
#       IF retries > 0
#         WAIT for (1.5^retries * 500) milliseconds +- some jitter
#
#       status = makeCallToForm3(timeout:DEFAULT_TIMEOUT)
#
#       IF status = SUCCESS (2xx) or CONFLICT (409)
#           retry = false
#       ELSE IF status = THROTTLED (429) # You have reached your request limit and are being throttled
#           retry = true
#      ELSE IF status >= 500 # A temporary issue has occurred, all requests are idempotent and safe to retry
#           retry = true
#       ELSE # Another http response such as 400 bad request, client must fix request before retrying
#           retry = false
#       END IF
#     CATCH EXCEPTION
#       retry = true # connection timeout, connection dropped etc...
#     END TRY
#
#     retries = retries + 1
# WHILE (retry AND (retries <= MAX_RETRIES))

*/
