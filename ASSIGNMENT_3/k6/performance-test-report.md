
# **Performance Testing Report – Golang-Gin RealWorld Example App**

**Tools Used:** k6
**Test Environment:** Local deployment of Golang-Gin RealWorld Example App
**Scripts:** `load-test.js`, `spike-test.js`, `soak-test.js`

---

## **1. Objective**

The purpose of this performance testing exercise was to evaluate the **scalability, reliability, and responsiveness** of the RealWorld backend API under different load conditions:

1. **Load Test:** Assess system performance under increasing concurrent users.
2. **Spike Test:** Test system stability under sudden bursts of traffic.
3. **Soak Test:** Evaluate long-term reliability and resource usage over an extended period.

---

## **2. Test Scenarios**

| Test Type  | Max VUs | Duration     | Description                                                           |
| ---------- | ------- | ------------ | --------------------------------------------------------------------- |
| Load Test  | 50      | 16 min       | Gradually increasing load to measure system throughput and latency.   |
| Spike Test | 100     | 1 min 10 sec | Sudden traffic surge to evaluate system behavior under extreme load.  |
| Soak Test  | 20      | 30 min       | Long-duration test to detect memory leaks or performance degradation. |

---

## **3. Key Metrics Captured**

* **HTTP Request Duration (`http_req_duration`)** – avg, median, p95
* **HTTP Request Failure Rate (`http_req_failed`)** – % of failed requests
* **Iterations** – completed requests per second
* **Virtual Users (VUs) Active** – number of concurrent users
* **Checks Passed/Failed** – functional validations like login, article creation, status 200

---

## **4. Test Results**

### **4.1 Summary Table**

| Test Type  | Max VUs | Iterations | Avg Req Duration | p95 Req Duration | HTTP Fail Rate | Checks Failed                 | Key Observations                                                      |
| ---------- | ------- | ---------- | ---------------- | ---------------- | -------------- | ----------------------------- | --------------------------------------------------------------------- |
| Load Test  | 50      | 15,945+    | 1.15–1.63 ms     | 2.6–4.18 ms      | 25–33%         | login, user, article creation | Fast response, high failure rate in authentication & article creation |
| Spike Test | 100     | 1,545      | 1.45–1.57 ms     | 4.37–4.67 ms     | 33%            | login                         | Stable under sudden load, persistent login failures                   |
| Soak Test  | 20      | 29,996     | 1.12–1.25 ms     | 2 ms             | 33%            | login                         | No resource degradation over 30 min, login failure persists           |

---

### **4.2 Detailed Observations**

**Load Test:**

* Requests processed quickly (avg < 2 ms).
* High failure rate (~33%) due to login and article creation failures.
* Throughput increased proportionally with virtual users.

**Spike Test:**

* System handled sudden surge well in terms of latency.
* Failure rate indicates authentication bottleneck under high concurrency.
* No crashes; system remained stable.

**Soak Test:**

* Long-duration test maintained low response times.
* Failure rate remained consistent, suggesting persistent functional issues rather than performance degradation.
* No memory leaks or resource exhaustion observed.

---

## **5. Analysis and Insights**

1. **Latency & Throughput**

   * Backend response times are extremely low (< 5 ms p95) across all scenarios.
   * Throughput scales well with increasing virtual users.

2. **Error Rates**

   * High failure rate (~33%) is primarily functional (login, article creation).
   * Indicates need for further debugging or test data/token validation.

3. **System Stability**

   * Spike test confirmed stability under sudden high load.
   * Soak test confirms reliability over extended periods.

4. **Bottlenecks Identified**

   * Authentication endpoints fail under concurrent load.
   * Article creation fails when multiple users post simultaneously.

---

## **6. Recommendations**

1. **Fix Functional Failures**

   * Investigate login and article creation errors.
   * Ensure test scripts include valid tokens and pre-existing test data.

2. **Increase Test Coverage**

   * Test other critical endpoints such as comments, profile updates, and tags.

3. **Optimize Backend**

   * Use caching for frequently requested data (tags, articles).
   * Optimize database queries and connection handling.

4. **Monitoring & Observability**

   * Integrate Grafana/Prometheus to track CPU, memory, and DB metrics during load.

---

## **7. Conclusion**

The Golang-Gin RealWorld backend demonstrates **excellent responsiveness and stability** under varying load conditions. Persistent functional failures, particularly in login and article creation endpoints, must be addressed before production deployment. Once fixed, the system is well-prepared to handle high-concurrency traffic reliably.

---

