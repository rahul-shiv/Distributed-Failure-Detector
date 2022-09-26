"""
Script for integration testing distributed grep.
"""


import os
import re
import time

CLIENT_FILE = "client.go"
OPTIONS = "c"
REGEX_OPTION = "cH"
CLIENT_DIR = os.getcwd() + "/../client/"
LOG_DIR = "/mp1/logs/"
TEST_CLIENT_DIR = os.getcwd() + "/../test_client/"

'''dictionary with expected values for queries against generated logs
'''
RIGHT_ANSWERS = {
"quo":{
"tests/test1.log":407100,
"tests/test10.log":271400,
"tests/test2.log":271400,
"tests/test3.log":135700,
"tests/test4.log":0,
"tests/test5.log":135700,
"tests/test6.log":135700,
"tests/test7.log":135700,
"tests/test8.log":271400,
"tests/test9.log":407100},

"cupiditate":{
"tests/test1.log":0,
"tests/test10.log":0,
"tests/test2.log":135700,
"tests/test3.log":0,
"tests/test4.log":0,
"tests/test5.log":0,
"tests/test6.log":0,
"tests/test7.log":0,
"tests/test8.log":0,
"tests/test9.log":0},

"rshiv":{
"tests/test1.log":0,
"tests/test10.log":0,
"tests/test2.log":0,
"tests/test3.log":0,
"tests/test4.log":0,
"tests/test5.log":0,
"tests/test6.log":0,
"tests/test7.log":0,
"tests/test8.log":0,
"tests/test9.log":0},

"distributed":{
"tests/test7.log":271400,
"tests/test10.log":135700,
"tests/test3.log":135700,
"tests/test5.log":407100,
"tests/test9.log":271400,
"tests/test4.log":678500,
"tests/test6.log":678500,
"tests/test1.log":542800,
"tests/test2.log":542800,
"tests/test8.log":135700},

"^Qui":{
"tests/test3.log":0,
"tests/test9.log":0,
"tests/test7.log":0,
"tests/test10.log":135700,
"tests/test4.log":0,
"tests/test5.log":0,
"tests/test6.log":0,
"tests/test1.log":135700,
"tests/test8.log":135700,
"tests/test2.log":0},

"'\sdolorem'":{
"tests/test5.log":0,
"tests/test9.log":0,
"tests/test4.log":0,
"tests/test6.log":0,
"tests/test8.log":0,
"tests/test2.log":0,
"tests/test3.log":135700,
"tests/test10.log":135700,
"tests/test7.log":135700,
"tests/test1.log":271400}
}

'''function to run the distributed grep client program
'''
def runDistGrep(options,pattern,file_name):
    os.chdir(CLIENT_DIR)
    cmd = "go run {0} {1} {2} {3}".format(CLIENT_FILE,options,pattern,file_name)
    print(cmd)
    out = os.popen(cmd).read()
    return out

'''function to parse the output into a dictionary
'''
def parse_output(str):
    parsed_out = {}
    result = re.findall("(.*):([0-9]+)",str)
    for file,count in result:
        parsed_out[file]=int(count)
    return parsed_out

'''function to validate if the parsed output matches
    the known counts for the pattern
'''
def assert_test(out,pattern,testname):
    parsed_output = parse_output(out)
    for file,count in RIGHT_ANSWERS[pattern].items():
        if file not in parsed_output or parsed_output[file]!=count:
            print("TEST {} FAILED".format(testname))
            return
    if len(parsed_output)!=len(RIGHT_ANSWERS[pattern]):
        print("TEST {} FAILED".format(testname))
    else:
        print("TEST {} PASSED".format(testname))

''' negative test case function to check if the expected output is received
    when one of the worker nodes goes down
'''
def assert_negative(out,testname,killed_process,pattern):
    parsed_output = parse_output(out)
    for file,count in RIGHT_ANSWERS[pattern].items():
        if file == "tests/test{}.log".format(killed_process):
            if file in parsed_output:
                print("lame {} TEST {} FAILED".format(file,testname))
                return
        elif file not in parsed_output or parsed_output[file]!=count:
            print("lamest {} TEST {} FAILED".format(file,testname))
            return
    print("TEST {} PASSED".format(testname))

'''function to create test logs
'''
def create_logs():
    os.chdir(TEST_CLIENT_DIR)
    cmd = "go run test_client.go create"
    os.system(cmd)

'''function to delete test logs
'''
def delete_logs():
    os.chdir(TEST_CLIENT_DIR)
    cmd = "go run test_client.go delete"
    os.system(cmd)

def TestDistributedGrepInfrequentWord():
    create_logs()
    pattern = "cupiditate"
    file_name = "tests/'*'"
    out = runDistGrep(OPTIONS,pattern,file_name)
    print(out)
    assert_test(out,pattern,"TestDistributedGrepInfrequentWord")
    delete_logs()

def TestDistributedGrepFrequentWord():
    create_logs()
    pattern = "distributed"
    file_name = "tests/'*'"
    os.chdir(CLIENT_DIR)
    out = runDistGrep(OPTIONS,pattern,file_name)
    print(out)
    assert_test(out,pattern,"TestDistributedGrepInfrequentWord")
    delete_logs()

def TestDistributedGrepRegex():
    create_logs()
    pattern = "'\sdolorem'"
    file_name = "tests/'*'"
    os.chdir(CLIENT_DIR)
    out = runDistGrep(REGEX_OPTION,pattern,file_name)
    print(out)
    assert_test(out,pattern,"TestDistributedGrepRegex")
    delete_logs()

def TestDistributedGrepNonExistentWord():
    create_logs()
    pattern = "rshiv"
    file_name = "tests/'*'"
    os.chdir(CLIENT_DIR)
    out = runDistGrep(OPTIONS,pattern,file_name)
    print(out)
    assert_test(out,pattern,"TestDistributedGrepInfrequentWord")
    delete_logs()

def TestDistributedGrepKillRandomWorker():
    create_logs()
    pattern = "quo"
    file_name = "tests/'*'"
    os.chdir(CLIENT_DIR)
    killed_process = "1"
    os.system("ssh fa22-cs425-030{} 'pkill worker_server'".format(killed_process))
    out = runDistGrep(OPTIONS,pattern,file_name)
    print(out)
    assert_negative(out,"TestDistributedGrepKillRandomWorker",killed_process,pattern)
    os.system("ssh fa22-cs425-030{} 'cd /mp1/cs425-mp1 ; go run worker/worker_server.go > /dev/null 2>&1 &'".format(killed_process))
    time.sleep(5)
    delete_logs()

TestDistributedGrepInfrequentWord()
TestDistributedGrepFrequentWord()
TestDistributedGrepRegex()
TestDistributedGrepNonExistentWord()
TestDistributedGrepKillRandomWorker()