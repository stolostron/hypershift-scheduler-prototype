# hypershift-scheduler-prototype
Prototyping scheduling hosted cluster placement to a set of hosting service clusters based on the current load

---

What we know:
    
* `AddOnPlacementScore`, mentioned in [extensible scheduling](https://open-cluster-management.io/concepts/placement/#extensible-scheduling) was added in OCM v0.6.0
* As of June 29, Placement API [supports resource based scheduling](https://github.com/open-cluster-management-io/community/issues/52#issuecomment-1170703250)
* Placement can be triggered by clusterclaim values.

What we're unsure of:

* What is cluster/workload-affinity/anti-affinity? Is it relevant to us?
* What is spreading policy?
* Is [this document](https://docs.google.com/document/d/1aaRSnyChczgJvejzug-fkY2eYhhQWfa2CPmqRq5s-SA/edit?pli=1#) relevant to the current project? If so, how? Is it out of date (seems to last be updated >year ago)? 
* Do we need a new weighting scheme?
* What does "scheduling" entail? Does the placement of hosted clusters change periodically depending on a set of criteria? If not, why is it referred to as *"scheduling"*?

* It seems like a lot of scheduling placement work was done recently. Is the goal of this project simply to tie it all together/make it automatic? [This document](https://docs.google.com/document/d/1066z0A7ZcCefsfDtbNKGGVSy2BPK5ZdZ-gHj5MxS-DY/edit?pli=1#heading=h.vrm2bomkuj48) includes a draft for the enhancement: scheduling based on cluster capacity/resource usage. What's the difference between this and "scheduling hosted cluster placement to a set of hosting service clusters based on the current load"?


What's been done:

* `AddOnPlacementScore` API, mentioned in [extensible scheduling](https://open-cluster-management.io/concepts/placement/#extensible-scheduling) was added in OCM v0.6.0. The accepted changes are outlined in the [Proposal](https://docs.google.com/document/d/1066z0A7ZcCefsfDtbNKGGVSy2BPK5ZdZ-gHj5MxS-DY/edit?pli=1#heading=h.yo0zi8mphpx7), [Examples](https://docs.google.com/document/d/1066z0A7ZcCefsfDtbNKGGVSy2BPK5ZdZ-gHj5MxS-DY/edit?pli=1#heading=h.c5bo6qw6z28h) and [Design Details](https://docs.google.com/document/d/1066z0A7ZcCefsfDtbNKGGVSy2BPK5ZdZ-gHj5MxS-DY/edit?pli=1#heading=h.gedsgt61to7w) outlined in the document mentioned in [this issue](https://github.com/open-cluster-management-io/community/issues/52#). 

What needs to be done:

* From what I understand, we need to (schedule?) placing hosted clusters on a set of hosting clusters based on the current load i.e. # of hosted clusters on the hosting cluster.


Conclusion: We need to (schedule?) placing hosted clusters on a set of hosting clusters based on the current load. However, I'm failing to see the difference between this and [previous work](https://github.com/open-cluster-management-io/community/issues/52#). Is something different needed for Hypershift? Was just the API implemented, but now I need to automate the process? If it's simply automation, why is the word "prototype" used when that implies the need to develop a method to select hosting clusters?


Some questions that arose this week when working on https://issues.redhat.com/browse/ACM-1540 

Questions: 
1. From my understanding, we need to (schedule?) placing hosted clusters on a set of hosting clusters based on the current load. However, I'm failing to see the difference between this and [previous work](https://github.com/open-cluster-management-io/community/issues/52#). Is something different needed for Hypershift? Was just the API implemented, but now I need to automate the process? If it's simply automation, why is the word "prototype" used when that implies the need to develop a method to select hosting clusters? What am I missing?
2. Is [this document](https://docs.google.com/document/d/1aaRSnyChczgJvejzug-fkY2eYhhQWfa2CPmqRq5s-SA/edit?pli=1#) relevant to the current project? If so, in what sense? Is it out of date (seems to last be updated more than a year ago)? 


---

SCRUM notes:

* 